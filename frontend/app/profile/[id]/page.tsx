import { Avatar } from "@/components/avatar";
import ModeToggle from "@/components/mode-toggle";
import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { shortWord } from "@/lib/utils";
import { Metadata } from "next";
import Head from "next/head";
import { useMemo } from "react";

// const chat:{
//   href:string,
//   user:string,
//   message:string
// }[] = [
//   {user: '@nursyah', message: 'hi mom', href:'/chat/nursyah'}
// ]

// const item = () => {
//   const res = []
//   for(let i=0; i < 100; i++){
//     res.push({user: '@nursyah', message: 'hi mom', href:'/chat/nursyah'})
//   }
//   return res
// }

// export const metadata: Metadata = {
//   title: `chat | ${process.env.APP}`,
// };

export async function generateMetadata({
  params,
}: {
  params: { id: string };
}): Promise<Metadata> {
  // read route params
  const id = params.id;

  // console.log(params);
  return {
    title: `profile - ${params.id} | ${process.env.APP}`,
  };
}

export default function Home({ params }: { params: { id: string } }) {
  return (
    <main className="p-4 w-full flex justify-center">
      <div className="flex max-w-2xl w-full p-2 flex-col">
        {/* navigation */}
        <nav className="flex justify-between w-full items-center border-b pb-2">
          <a href="/">Chat App</a>
        </nav>

        {/* profile and username */}
        <div className="flex justify-center items-center h-[80vh] flex-col">
          <div className="gap-y-4 flex flex-col w-full max-w-md items-center">
            <Avatar
              big
              src="https://priv.au/image_proxy?url=https%3A%2F%2Fencrypted-tbn0.gstatic.com%2Fimages%3Fq%3Dtbn%3AANd9GcS5V4CfhwPBr_gsiA7Fr4-73O542ZUWXu2zhkoY2QABQyh9lN8%26s&h=6338b107704da38caa66d438f4e90d1907314946d9d376dc6808f5eae49c1691"
            />

            <h1 className="font-bold text-2xl">@{params.id}</h1>
          </div>
        </div>
      </div>
    </main>
  );
}
