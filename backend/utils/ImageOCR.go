package utils

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

type EntityAnnotation = visionpb.EntityAnnotation

func ocr(filepath string) []*EntityAnnotation {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return nil
	}
	defer client.Close()

	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return nil
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Printf("Failed to create image: %v", err)
		return nil
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 5)
	if err != nil {
		log.Printf("Failed to detect texts: %v", err)
		return nil
	}

	// 最初は全体をテキストにした結果が入るので飛ばす
	annotations = annotations[1:]
	return annotations
}

func getTotalPayment(annotations []*EntityAnnotation, target string, threshold float64) (int, error) {
	yInterval := [2]int32{-1, -1}
	var xBoundary int32 = -1
	for _, annotation := range annotations {
		if strings.Contains(annotation.Description, target) {
			yInterval[0] = annotation.BoundingPoly.Vertices[0].Y
			yInterval[1] = annotation.BoundingPoly.Vertices[3].Y
			xBoundary = annotation.BoundingPoly.Vertices[1].X
		}
	}
	paymentText := ""
	for _, annotations := range annotations {
		targetYInterval := [2]int32{annotations.BoundingPoly.Vertices[0].Y, annotations.BoundingPoly.Vertices[3].Y}
		targetX := annotations.BoundingPoly.Vertices[1].X
		if math.Abs((float64(targetYInterval[0]-yInterval[0]))) < threshold && math.Abs(float64(targetYInterval[1]-yInterval[1])) < threshold && targetX > xBoundary {
			paymentText += annotations.Description
		}
	}
	re := regexp.MustCompile(`[¥,.]`)
	paymentText = re.ReplaceAllString(paymentText, "")
	return strconv.Atoi(paymentText)
}

func fileToTexts(filepath string) []string {
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	texts := []string{}
	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	return texts
}

// 各商品の値引額がよくわからない場合
// 経験値として単純な値引額の合計を返す
func annotationsToPoint(annotations []*EntityAnnotation, threshold float64) int {
	discountSum := 0

	for _, annotation := range annotations {
		if strings.HasPrefix(annotation.Description, "-") {
			num, err := strconv.Atoi(annotation.Description[1:])
			if err == nil {
				discountSum += num
			}
		}
	}
	// total, err := getTotalPayment(annotations, "合計", threshold)
	// fmt.Println(total, discountSum)
	// if err != nil {
	// 	log.Fatalf("Failed to get total payment: %v", err)
	// }

	return discountSum * 100
}

// 各商品の値引き額がわかる場合
// 割引率の合計(経験値)と割引商品の数（貢献度）を返す
func calculatePoint(annotations []*EntityAnnotation, threshold float64) (int, int) {
	pts := 0
	cnt := 0
	// 左上のポイントのY座標で昇順ソート
	sort.Slice(annotations, func(i, j int) bool {
		return annotations[i].BoundingPoly.Vertices[0].Y > annotations[j].BoundingPoly.Vertices[0].Y
	})
	for idx, annotation := range annotations {
		if strings.HasPrefix(annotation.Description, "-") {
			discountValue, err := strconv.Atoi(annotation.Description[1:])
			xPosition := annotation.BoundingPoly.Vertices[0].X
			cnt++
			// -(数字) というものが見つかったら
			// その左上のポイントのY座標より小さいY座標を持つ数字の中で最も大きいY座標を持つものを見つける
			// すなわちそれより後に出てくる数字のうち最初に出てくるものをとる
			if err == nil {
				for _, underAnnotation := range annotations[idx:] {
					underXPosition := underAnnotation.BoundingPoly.Vertices[0].X
					listPrice, err := strconv.Atoi(underAnnotation.Description)
					if err == nil && listPrice > discountValue && math.Abs(float64(underXPosition-xPosition)) < threshold {
						//fmt.Println(listPrice, discountValue)
						discountRate := 100 * discountValue / listPrice
						pts += discountRate
						break
					}
				}
			}
		}
	}
	return pts, cnt * 200
}

// エンコード
func encode(fileName string) string {

	file, _ := os.Open(fileName)
	defer file.Close()

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

// デコードを行い　一時的に作業ディレクトリに再画像化したものを保存
func decode(str string) string {
	data, _ := base64.StdEncoding.DecodeString(str) //[]byte
	currentDir, _ := os.Getwd()
	decodePath := currentDir + "/decode" + timeToString(time.Now())[:11] + ".png" //画像保存用パス 実行時の時間を元に作成する
	file, _ := os.Create(decodePath)

	defer file.Close()

	file.Write(data)

	return decodePath
}

func timeToString(t time.Time) string { //時間を文字列に変換するための関数
	str := t.Format("2006-01-02T15:04:05Z07:00")
	return str
}

func removeFile(filePath string) bool { //画像ファイル消去 正しく消去できなかった場合に以降の処理が進まないように
	r := os.Remove(filePath)
	if r != nil {
		fmt.Println("画像の保存に失敗しました")
		return false
	} else {
		return true
	}
}

func ForCharacterController(receipt string) (int, int, error) {
	decodePath := decode(receipt)
	annotations := ocr(decodePath)
	if annotations == nil {
		return 0, 0, fmt.Errorf("failed to ocr")
	}
	remove := removeFile(decodePath)
	if remove {
		exp, cnt := calculatePoint(annotations, 10)
		return exp, cnt, nil
	}
	return 0, 0, fmt.Errorf("failed to calculate point")
}

func main() {
	encData := encode("/Users/Hiroki Yoshida/go test/スクリーンショット 2023-08-22 113438.png") //入力画像
	decodePath := decode(encData)                                                      //base64でencodeされたデータに対してdecodeを行い結果を保存→ ファイルパスを返す
	annotations := ocr(decodePath)
	remove := removeFile(decodePath)
	if remove {
		fmt.Println(calculatePoint(annotations, 10))
	}
}
