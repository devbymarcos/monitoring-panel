const API_URL = import.meta.env.VITE_API_URL || "";

export async function getStatusHttp() {
  console.time("getStatusHttp request");
  const response = await fetch(`${API_URL}/api/status`);
  if (!response.ok) {
    throw new Error("Failed to fetch status");
  }
  const data = await response.json();
  console.timeEnd("getStatusHttp request");
  console.log("getStatusHttp response:", data);
  return data;
}
