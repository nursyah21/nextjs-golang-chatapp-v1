"use client";
import { Avatar } from "@/components/avatar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Edit, SendHorizonal } from "lucide-react";
import { useForm, SubmitHandler } from "react-hook-form";

type Inputs = {
  username: string;
};

export default function EditProfile() {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Inputs>({
    defaultValues: {
      username: "nursyah",
    },
  });
  const onSubmit: SubmitHandler<Inputs> = (data) => console.log(data);

  // console.log(watch("username"));

  const imageUpload = (e:any) => {
    console.log(e)
  }

  return (
    <div className="flex justify-center items-center h-[80vh] flex-col">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="gap-y-4 flex flex-col w-full max-w-md"
      >
        <Label htmlFor="profile">
          <Avatar className="cursor-pointer"
            big
            src="https://priv.au/image_proxy?url=https%3A%2F%2Fencrypted-tbn0.gstatic.com%2Fimages%3Fq%3Dtbn%3AANd9GcS5V4CfhwPBr_gsiA7Fr4-73O542ZUWXu2zhkoY2QABQyh9lN8%26s&h=6338b107704da38caa66d438f4e90d1907314946d9d376dc6808f5eae49c1691"
          />
        </Label>
        <input id="profile" type="file" accept="image/*" className="hidden" onChange={e=>imageUpload(e.target.files)} />


        <h1 className="font-bold text-2xl">@nursyah</h1>

        <div className="flex flex-col w-full space-y-2">
          <Label htmlFor="username">Username</Label>
          <div className="flex">
            <Input id="username" placeholder="username" {...register("username")} />
            <Button variant={"ghost"} type="submit">
              <Edit />
            </Button>
          </div>
        </div>
        
        <Button>Logout</Button>
      </form>
    </div>
  );
}
