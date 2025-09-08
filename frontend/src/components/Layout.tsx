import { Outlet } from "react-router-dom";
import { AppSidebar } from "./AppSidebar";
import HeaderNav from "./HeaderNav";


export default function Layout() {
  return (
    <div className="min-h-screen  bg-dbmcustom-bg text-dbmcustom-text">
      <AppSidebar />
      <main className="ml-[250px] p-2">
        <HeaderNav/>
        <Outlet />
      </main>
    </div>
  );
}
