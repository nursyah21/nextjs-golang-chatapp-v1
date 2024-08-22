"use client";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { SendHorizonal } from "lucide-react";
import { useForm, SubmitHandler } from "react-hook-form";

type Inputs = {
  message: string;
};

export default function ChatList() {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Inputs>();
  const onSubmit: SubmitHandler<Inputs> = (data) => console.log(data);
  // console.log(watch("message"));

  return (
    <>
      {
        items().map((e,idx)=>(
          <Chat key={idx} message={e.message} received={e.received} />
        ))
      }
    </>
  );
}

const items = () => {
  const res = []
  for(var i=0; i < 10; i++){
    res.push({message:"hehe kenapa bang ada masalaha apaa sdasdjnwiwnejicqwjciwenciwqvjiwenqjinjciwnevjiwnqjivnwejivnwidvndwjiqvn",
      received: i % 2 == 0 ? false : true})
  }
  return res
}

const Chat = (props: { message: string, received?: boolean }) => (
  <div className={!props.received ? "flex justify-end" : ''}>
    <Card className="my-2 p-2 max-w-sm rounded-2xl break-words">
      {props.message}
    </Card>
  </div>
);
