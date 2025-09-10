import { useState } from "react";
import { TvMinimal } from "lucide-react";

const MonitoringItem = ({ mdata }) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="bg-dbmcustom-bg px-4 py-2 rounded ">
      {/* Cabeçalho sempre visível */}
      <div
        className="flex items-center gap-2 cursor-pointer"
        onClick={() => setIsOpen(!isOpen)}
      >
        <div className="flex-1 flex items-center gap-4 text-sm">
          {mdata.status === "online" && (
            <span className="relative flex size-3">
              <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-green-400 opacity-75"></span>
              <span className="relative inline-flex size-3 rounded-full bg-green-500"></span>
            </span>
          )}
          {mdata.status === "offline" && (
            <span className="relative flex size-3">
              <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-red-400 opacity-75"></span>
              <span className="relative inline-flex size-3 rounded-full bg-red-500"></span>
            </span>
          )}

          <TvMinimal color="#fa7b05" />

          <div>
            <p>{mdata.name}</p>
            <span className="text-dbmcustom-textMuted text-[11px]">
              {mdata.ip}
            </span>
          </div>
        </div>

        {/* Ícone indicador */}
        <span className="text-xs text-dbmcustom-textMuted">
          {isOpen ? "▲" : "▼"}
        </span>
      </div>

      {/* Conteúdo expandível */}
      <div
        className={`overflow-hidden transition-all duration-300 text-sm text-dbmcustom-textMuted ${
          isOpen ? "max-h-40 mt-2" : "max-h-0"
        }`}
      >
        <p>
          <strong>Status:</strong> {mdata.status}
        </p>
        <p>
          <strong>Última mudança:</strong> {mdata.ultima_mudanca}
        </p>
        <p>
          <strong>IP:</strong> {mdata.ip}
        </p>
      </div>
    </div>
  );
};

export default MonitoringItem;
