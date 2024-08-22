"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { SendHorizonal } from "lucide-react";
import { useForm, SubmitHandler } from "react-hook-form";

type Inputs = {
  message: string;
};

export default function Message() {
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
      <form onSubmit={handleSubmit(onSubmit)} className="">
        {/* message */}
        <div className="fixed bottom-0 py-4 w-full max-w-2xl bg-[hsl(var(--background))]">
          <div className="flex w-full space-x-2">
            <Input placeholder="message" {...register("message")} />
            <Button variant={"ghost"} type="submit">
              <SendHorizonal />
            </Button>
          </div>
        </div>
      </form>
    </>
  );
}
