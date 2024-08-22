"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useForm, SubmitHandler } from "react-hook-form";

type Inputs = {
  username: string;
  password: string;
  confirmpassword: string;
};

export default function Form() {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Inputs>();
  const onSubmit: SubmitHandler<Inputs> = (data) => console.log(data);
  // console.log(watch("username"));

  return (
    <>
      <div className="flex justify-center items-center h-[80vh] flex-col">
        <form
          className="gap-y-4 flex flex-col w-full max-w-md"
          onSubmit={handleSubmit(onSubmit)}
        >
          <h1 className="font-bold text-2xl">Register</h1>

          <div>
            <Label htmlFor="username">Username</Label>
            <Input
              id="username"
              placeholder="username"
              {...register("username")}
            />
          </div>
          <div>
            <Label htmlFor="password">Password</Label>
            <Input
              id="password"
              placeholder="password"
              {...register("password")}
            />
          </div>
          <div>
            <Label htmlFor="confirm-password">Confirm Password</Label>
            <Input
              id="confirm-password"
              placeholder="confirm password"
              {...register("confirmpassword")}
            />
          </div>

          <Button type="submit" className="w-full">
            Submit
          </Button>

          <a className="text-xs hover:underline" href="/login">
            Login
          </a>
        </form>
      </div>
    </>
  );
}
