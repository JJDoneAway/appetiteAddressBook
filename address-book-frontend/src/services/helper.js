export function validateEmail(email) {
  const re =
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@(([\w-]+\.)+[a-zA-Z]{2,})$/;
  return re.test(email);
}

export function validatePhone(phone) {
  var re = /\+\d{5,50}/;
  return re.test(phone);
}

export function parseJwt(token) {
  const header = decodeJWT(token, 0);
  const payload = decodeJWT(token, 1);
  if (payload.exp !== undefined) {
    payload.exp = format_time(payload.exp);
    payload.iat = format_time(payload.iat);
  }
  return { header: header, payload: payload, signature: token.split(".")[2] };
}

function decodeJWT(token, section) {
  try {
    var base64Url = token.split(".")[section];
    var base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    var jsonPayload = decodeURIComponent(
      window
        .atob(base64)
        .split("")
        .map(function (c) {
          return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join("")
    );
    return JSON.parse(jsonPayload);
  } catch (err) {
    return JSON.parse(
      `{"error": "Not possible to parse the content", "error-message": "${err}"}`
    );
  }
}

function format_time(s) {
  const d = new Date(s * 1e3);
  return d.toLocaleDateString() + " / " + d.toLocaleTimeString();
}
