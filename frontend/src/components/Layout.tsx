import { Outlet } from "react-router-dom";
import { AppSidebar } from "./AppSidebar";

export default function Layout() {
  return (
    <div className="min-h-screen  bg-dbmcustom-bg text-dbmcustom-text">
      <AppSidebar />
      <main className="ml-[250px] p-2">
        <Outlet />
      </main>
    </div>
  );
}
