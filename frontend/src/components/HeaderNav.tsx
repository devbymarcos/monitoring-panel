import { Menu } from "lucide-react";
import React from "react";

const HeaderNav = () => {
  return (
    <header className="p-4">
      <button className="bg-orange-400 p-1 rounded-sm cursor-pointer">
        <Menu size={20} />
      </button>
    </header>
  );
};

export default HeaderNav;
