
conclusion
==========================

[copying interface values in go](https://www.ardanlabs.com/blog/2016/05/copying-interface-values-in-go.html)

I allowed myself to get confused about what Go would do when making a copy of an interface value that was storing a value and not a pointer.
For a minute I wondered if each copy of the interface value also created a copy of the value referenced by the interface.
Since we were “storing” a value and not a pointer.
But what we learned is that since an address is always stored, it is the address that is being copied and never the value itself.

* 接口值是一个两个字长度的数据结构
  - 指向内部表的指针, iTable包含所存储的值的类型信息.
  - 指向实体值的指针.

