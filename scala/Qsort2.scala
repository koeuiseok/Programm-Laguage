object Qsort {

// Imperative method
def qsort(arr: Array[Int]): Array[Int] = {
  def swap(i: Int, j: Int): Unit = {
    val tmp = arr(i)
    arr(i) = arr(j)
    arr(j) = tmp
  }

  def partition(left: Int, right: Int): Unit =
    if (right <= left ) ()
    else {
      val pivot = arr((left + right) / 2)
      var i = left
      var j = right

      while (i <= j) {
        while (arr(i) < pivot) i += 1
        while (arr(j) > pivot) j -= 1

        if (i <= j) {
          swap(i, j)

          i += 1
          j -= 1
        }
      }

    if (left < j) partition(left, j)
    if (j < right) partition(i, right)
  }

  partition(0, arr.length - 1)
  arr
}

// if-else statement
def qsort2(seq: Seq[Int]): Seq[Int] =
   if (seq.isEmpty) seq
   else {
      val (pivot, rest) = (seq.head, seq.tail)
      val (lt, ge) = rest partition { _ < pivot }
      qsort2(lt) ++ (pivot +: qsort2(ge))
   }

// pattern match
def qsort3(seq: Seq[Int]): Seq[Int] =
   seq match {
   case pivot +: rest =>
       val (lt, ge) = rest partition { _ < pivot }
       qsort3(lt) ++ (pivot +: qsort3(ge))
   case _ =>
       seq
   }

  def main(args: Array[String]) {
        println("Input array:")

			  val arr: Array[Int] = Array(2, 4, 1, 3, 7, 8, 6, 9, 5)
			  arr.foreach(ele => println(ele))
			  var arr2 = qsort2(arr)

        println("Sorted array:")
			  arr2.foreach(println(_))
  }
}
