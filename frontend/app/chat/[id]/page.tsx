import { Avatar } from "@/components/avatar";
import ModeToggle from "@/components/mode-toggle";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { shortWord } from "@/lib/utils";
import { SendHorizonal } from "lucide-react";
import { Metadata } from "next";
import Message from "./message";
import ChatList from "./chat-list";


export async function generateMetadata(
  {params}: {params:{id:string}},
): Promise<Metadata> {
  // read route params
  const id = params.id

  // console.log(params)
  return {
    title: `chat - ${params.id} | ${process.env.APP}`,
  }
}

export default function Home({ params }: { params: { id: string } }) {
  
  return (
    <main className="p-4 w-full flex justify-center">
      
      <div className="flex max-w-2xl w-full p-2 flex-col">
        {/* navigation */}
        <nav className="flex justify-between w-full items-center border-b pb-2">
          <a href="/">Chat App</a>
          <a href={"/profile/" + params.id}>
            <Avatar />
          </a>
        </nav>

        {/* all chat */}
        <ChatList />

        {/* message */}
        <Message />

      </div>
    </main>
  );
}
