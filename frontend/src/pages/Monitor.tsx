import MonitoringItem from "@/components/MonitoringItem";
import { getStatusHttp } from "@/http/getStatusHttp";
import { useEffect, useState } from "react";

export default function Monitor() {
  const [data, setData] = useState([]);
  const [countdown, setCountdown] = useState(10);
  const m = Array.from({ length: 12 });

  const runStatus = async () => {
    const response = await getStatusHttp();
    setData(response);
    console.log("Monitor data state updated:", response);
  };

  useEffect(() => {
    const intervalId = setInterval(() => {
      runStatus();
    }, 10000);
    // Limpa o intervalo quando o componente é desmontado
    return () => clearInterval(intervalId);
  }, []);

  useEffect(() => {
    const countdownId = setInterval(() => {
      setCountdown((prev) => (prev <= 1 ? 10 : prev - 1));
    }, 1000);
    return () => clearInterval(countdownId);
  }, []);

  return (
    <>
      <h1 className="text-2xl text-center font-bold">
        Painel de Monitoramento
      </h1>
      <p className="text-center text-sm mt-2">
        Próxima atualização em:{" "}
        <span className="text-orange-300 mr-1.5">{countdown}</span>segundos
      </p>
      <section className=" bg-dbmcustom-panel grid grid-cols-4 gap-3 rounded-lg p-4 mt-6 ">
        {data.map((monitorGroup, groupIndex) => (
          <div key={groupIndex} className="col-span-4">
            <h2 className="text-xl font-semibold mb-2">
              {monitorGroup.comunication_service}
            </h2>
            <div className="grid grid-cols-4 gap-3">
              {monitorGroup.itens.map((item, itemIndex) => (
                <MonitoringItem key={itemIndex} mdata={item} />
              ))}
            </div>
          </div>
        ))}
      </section>
    </>
  );
}
