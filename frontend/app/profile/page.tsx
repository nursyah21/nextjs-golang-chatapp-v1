import { Avatar } from "@/components/avatar";
import ModeToggle from "@/components/mode-toggle";
import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { shortWord } from "@/lib/utils";
import { Metadata } from "next";
import EditProfile from "./edit-profile";

export const metadata: Metadata = {
  title: `profile | ${process.env.APP}`,
};

export default function Home({ params }: { params: { id: string } }) {
  return (
    <main className="p-4 w-full flex justify-center">
      <div className="flex max-w-2xl w-full p-2 flex-col">
        {/* navigation */}
        <nav className="flex justify-between w-full items-center border-b pb-2">
          <a href="/">Chat App</a>
          {/* <a href={"/" + params.id}>
            <Avatar />
          </a> */}
        </nav>
        <EditProfile />
      </div>
    </main>
  );
}
