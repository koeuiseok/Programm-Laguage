object OrderedArrayLookup {
 
  def binarySearch(array: Array[Int], value: Int) = {
    def printArray(array : Array[Int], start : Int, end : Int) {
      println(array.slice(start, end)mkString(","))
    }
     
    def binarySearchRecursive(array: Array[Int], start: Int, end: Int, value: Int): Int = {
      println("start -> " + start + ", end -> " + end)
      printArray(array, start, end+1)
      var i = (start + end) / 2
      if(start > end) {
        println("Not Found")
        return -99999999
      }
       
      array(i) match {
        case x if (x == value) => x
        case x if (x > value) => binarySearchRecursive(array, start, i - 1, value)
        case x if (x < value) => binarySearchRecursive(array, i + 1, end, value)
      }
    }
    binarySearchRecursive(array, 0, array.size - 1, value)
  }
   
  def main(args: Array[String]) {
   println("Found -> " + binarySearch(1 until 30 toArray, 11))
  }
}