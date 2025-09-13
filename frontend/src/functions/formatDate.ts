export function formatDate(date: Date) {
  const formatDate = new Date(date);
  return formatDate.toLocaleString("pt-BR");
}
