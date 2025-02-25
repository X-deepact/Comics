export function authHeader() {
  const access_token = localStorage.getItem("access_token");
  return {
    Authorization: `Bearer ${access_token}`,
  };
}
