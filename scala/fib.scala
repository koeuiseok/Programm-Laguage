def fib(n:Int, a:Int = 0, b:Int = 1): List[Int] = {
    if(a > n) Nil
    else a +: fib(n, b, a+b)
}
print(fib(13).mkString(", "))