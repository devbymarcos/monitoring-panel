import { TvMinimal } from "lucide-react";


const MonitoringItem = () => {
  return (<div className="bg-dbmcustom-bg px-4 py-2 rounded flex items-center gap-2">
    <div className="flex-1 flex items-center gap-2 text-sm">
      <TvMinimal />
     Equipamento 4 - Recepção - IP:192.168.10.20
    </div>
      <span className="relative flex size-3">
      <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-green-400 opacity-75"></span>
      <span className="relative inline-flex size-3 rounded-full bg-green-500"></span>
    </span>
  </div>)
};

export default MonitoringItem;
