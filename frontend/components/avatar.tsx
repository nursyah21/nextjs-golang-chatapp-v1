import { Avatar as A, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import Image from 'next/image'

export function Avatar(props: {
  src?: string,
  className?: string
  big?: boolean
}) {

  return (
    <A className={props.big ? "w-24 h-24  sm:w-32 sm:h-32" : ""}>
      <AvatarImage {...props} alt="@shadcn" />
      <AvatarFallback>
      <Image
            src="/profile.png"
            alt="profile"          
            width={500}  
            height={500}  
        />
      </AvatarFallback>
    </A>
  );
}
