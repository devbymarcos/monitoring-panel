import MonitoringItem from "@/components/MonitoringItem";

export default function Monitor() {
  const m = Array.from({ length: 12 });

  return (
    <>
      <h1 className="text-2xl text-center font-bold">
        Painel de Monitoramento
      </h1>
      <section className=" bg-dbmcustom-panel rounded-lg p-4 mt-6 h-screen">
        {m.map((_,index) => (
          <MonitoringItem name={index} />
        ))}
      </section>
    </>
  );
}
