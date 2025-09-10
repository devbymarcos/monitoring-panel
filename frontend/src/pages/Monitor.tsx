import MonitoringItem from "@/components/MonitoringItem";
import { getStatusHttp } from "@/http/getStatushttp";
import { useEffect, useState } from "react";

export default function Monitor() {
  const [data,setData] = useState([]);
  const m = Array.from({ length: 12 });


const runStatus = async () => {
  const response = await getStatusHttp();
  setData(response);
  console.log("request",response);
}

useEffect(() => {
  runStatus();
}, []);

  return (
    <>
      <h1 className="text-2xl text-center font-bold">
        Painel de Monitoramento
      </h1>
      <section className=" bg-dbmcustom-panel grid grid-cols-4 gap-3 rounded-lg p-4 mt-6 ">
        {data.map((monitor, index) => (
          <MonitoringItem key={index} mdata={monitor} />
        ))}
      </section>
    </>
  );
}
