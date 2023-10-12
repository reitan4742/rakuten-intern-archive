import VueCookie from "vue-cookie";
import { useRouter } from "vue-router";

export function checkCookie() {
    const userCookie = VueCookie.get("userName");
    const Router = useRouter();
    if (userCookie === null) {
        Router.push("/login");
    } else {
        const exp = calcExpirationDate();
        VueCookie.set("userName", userCookie, {
            expires: exp
        });
    }
}

export function calcExpirationDate() {
    const expirationDate = new Date();
    expirationDate.setTime(expirationDate.getTime() + 3 * 60 * 1000)
    return expirationDate;
}

export function getCookie() {
    return VueCookie.get("userName");
}