//class vs object
//object
object IdFactory {
  private var counter = 0
  def create(): Int = {
    counter += 1
    counter
  }
}
val newId: Int = IdFactory.create()
println(newId) // 1
val newerId: Int = IdFactory.create()
println(newerId) // 2

//class
class IdFactory_class() {
  var counter = 0
  def create(): Int = {
    counter += 1
    counter
  }
}
val newclassId = new IdFactory_class()
println(newclassId.create()) 
val newclassId2 = new IdFactory_class()
println(newclassId2.create()) 
