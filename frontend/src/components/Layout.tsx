import { Outlet } from "react-router-dom";
import { AppSidebar } from "./AppSidebar";
import HeaderNav from "./HeaderNav";
import { useState } from "react";

export default function Layout() {
  const [onsideBar, SetOnSideBar] = useState(false);

  return (
    <div className="min-h-screen  bg-dbmcustom-bg text-dbmcustom-text">
      <AppSidebar onsideBar={onsideBar} />
      <main
        className={`${
          onsideBar == true ? "ml-0" : "ml-[250px]"
        } transition-all p-2`}
      >
        <HeaderNav onsideBar={onsideBar} setOnsideBar={SetOnSideBar} />
        <Outlet />
      </main>
    </div>
  );
}
