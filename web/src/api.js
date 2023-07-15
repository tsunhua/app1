import axios from 'axios';

const BASE_URL = `http://${window.location.host}`
const KEY_TOKEN = 'app1_token'

export async function login(name, password) {
  try {
    const response = await axios.post(BASE_URL + "/api/login", {
      name: name,
      password: password,
    }, {
      headers: {
        'Content-Type': 'application/json'
      }
    });
    setCookie(KEY_TOKEN, response.data.app1_token, 7)
    return true
  } catch (error) {
    console.error(error);
    return false
  }
}

export async function querySites() {
  try {
    const response = await axios.get(BASE_URL + "/api/sites");
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

export function openApp(path) {
  const url = `${BASE_URL}/${path}`;
  window.open(url);
}

// 设置 Cookie
function setCookie(name, value, days) {
  const date = new Date();
  date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
  const expires = "expires=" + date.toUTCString();
  document.cookie = name + "=" + value + ";" + expires + ";path=/";
}

// 获取 Cookie
function getCookie(name) {
  const cookieName = name + "=";
  const cookieArray = document.cookie.split(';');
  for (let i = 0; i < cookieArray.length; i++) {
    let cookie = cookieArray[i];
    while (cookie.charAt(0) == ' ') {
      cookie = cookie.substring(1);
    }
    if (cookie.indexOf(cookieName) == 0) {
      return cookie.substring(cookieName.length, cookie.length);
    }
  }
  return "";
}

// 删除 Cookie
function deleteCookie(name) {
  document.cookie = name + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}