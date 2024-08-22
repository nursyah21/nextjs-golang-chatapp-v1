import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function shortWord(word:string){
  return word.length > 50 ? word.slice(0,50)+'...' : word
}