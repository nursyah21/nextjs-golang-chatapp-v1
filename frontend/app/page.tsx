import { Avatar } from "@/components/avatar";
import ModeToggle from "@/components/mode-toggle";
import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { shortWord } from "@/lib/utils";
import { useMemo } from "react";

const item = () => {
  const res = [];
  for (let i = 0; i < 1; i++) {
    res.push({ user: "nursyah", message: "hi mom", href: "/chat/nursyah" });
  }
  return res;
};

export default function Home() {
  return (
    <main className="p-4 w-full flex justify-center">
      <div className="flex max-w-2xl w-full p-2 flex-col">
        {/* navigation */}
        <nav className="flex justify-between w-full items-center border-b pb-2">
          <a href="/">Chat App</a>
          <a href="/profile">
            <Avatar />
          </a>
        </nav>

        {/* search */}
        <div className="my-4">
          <Input name="search" placeholder="search" />
        </div>

        {/* all contact and chat */}
        <div className="gap-y-4 flex flex-col">
          {item().map((e, idx) => (
            <CardName
              key={idx}
              message={e.message}
              href={e.href}
              user={e.user}
            />
          ))}
        </div>
      </div>

      {/* toggle theme */}
      <div className="fixed sm:bottom-8 sm:right-8 bottom-4 right-4">
        <ModeToggle className="rounded-full" />
      </div>
    </main>
  );
}

const CardName = (props: { href: string; user: string; message: string }) => (
  <div>
    <Card className="p-2">
      <div className="flex gap-x-2 flex-wrap">
        <div>
          <a href={"/profile/" + props.user}>
            <Avatar />
          </a>
        </div>
        <a className="text-xs flex-1" href={props.href}>
            <div className="flex gap-x-2 justify-between">
              <h1 className="font-bold underline">@{props.user}</h1>
              <h1 className="text-xs text-slate-600 dark:text-slate-400 ">2012-12-01 08:12</h1>
            </div>
          <h1 className="mt-1">{shortWord(props.message)}</h1>
        </a>
      </div>
    </Card>
  </div>
);
