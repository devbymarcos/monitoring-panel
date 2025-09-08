import { TvMinimal } from "lucide-react";

const MonitoringItem = ({ name }) => {
  return (
    <div className="bg-dbmcustom-bg px-4 py-2 rounded flex items-center gap-2 mb-2">
      <div className="flex-1 flex items-center gap-4 text-sm">
        <span className="relative flex size-3">
          <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-green-400 opacity-75"></span>
          <span className="relative inline-flex size-3 rounded-full bg-green-500"></span>
        </span>
        <TvMinimal color="#fa7b05" />

        <div>
          <p>Equipamento {name}</p>
          <span className="text-dbmcustom-textMuted text-[11px]">
            IP:192.168.10.20
          </span>
        </div>
      </div>
    </div>
  );
};

export default MonitoringItem;
