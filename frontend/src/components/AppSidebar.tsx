import React from "react";
import { Link } from "react-router-dom";
import { Binoculars, ChartArea } from "lucide-react";

export const AppSidebar = () => {
  return (
    <div className="h-screen w-[250px] fixed bg-dbmcustom-bg2 text-dbmcustom-text">
      <div className="h-20 flex  p-4 border-b border-dbmcustom-border">
        <p className="text-dbmcustom-textMuted font-bold flex flex-row items-center gap-2">
          <ChartArea />
          Monitoring Panel
        </p>
      </div>
      <nav>
        <ul>
          <li>
            <Link
              className="p-4 flex flex-row items-center gap-2 hover:bg-dbmcustom-bg transition-colors duration-300"
              to="/monitor"
            >
              <Binoculars size={16} />
              Monitor
            </Link>
          </li>
        </ul>
      </nav>
    </div>
  );
};
