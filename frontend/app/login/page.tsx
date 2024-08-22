import { Avatar } from "@/components/avatar";
import ModeToggle from "@/components/mode-toggle";
import { Card } from "@/components/ui/card";
import { shortWord } from "@/lib/utils";
import { useForm, SubmitHandler } from "react-hook-form"
import Form from "./form";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: `login | ${process.env.APP}`,
}

export default function Home() {  

  return (
    <main className="p-4 w-full flex justify-center">
      <div className="flex max-w-2xl w-full p-2 flex-col">

        {/* navigation */}
        <nav className="flex justify-between w-full items-center border-b pb-2">
          <a href="/">
            <h1>Chat App</h1>
          </a>
        </nav>

        <Form />

      </div>

      {/* toggle theme */}
      <div className="fixed sm:bottom-8 sm:right-8 bottom-4 right-4">
        <ModeToggle className="rounded-full"/>
      </div>
    </main>
  );
}

