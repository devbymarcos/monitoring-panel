import { Link } from "react-router-dom";
import { Binoculars, ChartArea } from "lucide-react";

export const AppSidebar = ({ onsideBar }) => {
  return (
    <div
      className={`h-screen ${
        onsideBar == true ? "-left-[250px]" : "left-0"
      } fixed w-[250px] bg-dbmcustom-bg2 transition-all text-dbmcustom-text`}
    >
      <div className="h-20 flex  p-4 border-b border-dbmcustom-border">
        <p className="text-dbmcustom-textMuted font-bold flex flex-row items-center gap-2">
          <ChartArea />
          Monitor
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
              Log monitor
            </Link>
          </li>
        </ul>
      </nav>
    </div>
  );
};
