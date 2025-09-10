const API_URL = import.meta.env.VITE_API_URL || "";

export async function getStatusHttp() {
  const response = await fetch(`${API_URL}/api/status`);
    if (!response.ok) {
        throw new Error("Failed to fetch status");
    }
    return response.json();
}