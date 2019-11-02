composition
=========================

[Composition with Go](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html)

### summary
* Declare the set of behaviors as discrete interface types first. 
  Then think about how they can be composed into a larger set of behaviors.
* Make sure each function or method is very specific about the interface types they accept. 
  Only accept interface types for the behavior you are using in that function or method.
  This will help dictate the larger interface types that are required.
* Think about embedding as an inner and outer type relationship. 
  Remember that through inner type promotion, everything that is declared in the inner type is promoted to the outer type.
  However, the inner type value exists in and of itself as is always accessible based on the rules for exporting.
* Type embedding is not subtyping nor subclassing.
  Concrete type values represent a single type and can’t be assigned based on any embedded relationship.
* The compiler can arrange interface conversions between related interface values. 
  Interface conversion, at compile time, doesn’t care about concrete types.
  it knows what to do merely based on the interface types themselves, not the implementing concrete values they could contain.
