# Learn GO

This is me following Nana's tutorial at https://www.youtube.com/watch?v=yyUHQIec83I for making a cli application with GO

- use _ to separate words in filenames.
- aliasing imports :: import fm "fmt"
- A constant’s value should be known at compile time according to the design principles where the function’s value is computed at run time.
     const x = getVal()
- Numeric constants have no size or sign. They can be of arbitrarily high precision and do not overflow:
     const x = 0.693147180559945309417232121458176568075500134360255254120680009
- unicode package has unicode.IsLetter(ch) unicode.IsDigit(ch) unicode.IsSpace(ch)
- strconv.Atoi(s string) (i int, err error) converts to an int
- strconv.Itoa() converts int to string
- var intP *int to declare a variable which will store location of an int. the later : intP = &someintvar
-  cannot take the address of a literal or a constant
    func main(){
    const i = 5
    ptr1 := &i //error: cannot take the address of i
    ptr2 := &10 //error: cannot take the address of 10
    }
- if initialization; condition  {} if supports initialization(good for multiple ret type)
- switch in go supports conditions and therefore need not be centered arounnd a variable :
    switch {
    case i < 0:
        f1() // function call
    case i == 0:
        f2() // function call
    case i > 0:
        f3() // function call
    }
- switch initialization; { switch supports ini. which can be later used
- multi updates in for ini.  for i, j := 0, N; i < j; i, j = i+1, j-1 { }
- to iterate characters in strings : for pos, char := range str {
- maybe don't use goto but here's the syntax :
    func main() {
        i:=0
        HERE:           // adding a label for a location
        fmt.Printf("%d",i)
        i++
        if i==5 {
            return
        }
        goto HERE      // goto HERE 
    }
- default values if not initialized yet is 0 for all numeric data types and byte too. for complex nums 0 + 0i, for strings ""
- function overloading is not supported in go
- Named result parameters are automatically initialized to their default-value of data type, and once they receive their value, a simple (empty) return    statement is sufficient.
        func SumProductDiff(i, j int)(s int, p int, d int) {   // named version
            s = i + j
            p = i * j
            d = i - j
            return
        }
- variable number of params (like ...rest param in node) func Greeting(prefix int, who ...string) :: who is created as a slice

- DEFER keyword
    func main() {
        Function1()
    }

    func Function1() {
        fmt.Printf("In Function1 at the top\n")
        defer Function2() // function deferred, will be executed after Function1 is done.
        fmt.Printf("In Function1 at the bottom!\n")
    }

    func Function2() {
        fmt.Printf("Function2: Deferred until the end of the calling function!")
    }

    DEFER works LIFO . If 5 things are deferred, the later defered stuff will run first.

    but any pre function whose value is passed to the outermost function that is deferred is calculated first :

    func deferred (x string) {
	fmt.Printf("%v lol\n", x)
    }

    func pre () string {
        fmt.Printf("pre\n")
        return "pre"
    }

    func play () {

        defer deferred(pre())
        fmt.Printf("play\n")

    }

    outputs : pre
              play
              pre lol

- built-in functions :
    - len() gives length 
    - cap() gives max size the slice can be increased to.
    - new(T) allocates zeroed storage for a new item of type T and returns its address. v := new(int) // v has type *int
    - append(to, ...values) used with slices. append method increases capacity of dest (if slice) if need be
    - copy(dst, src) copies a elements to b. returns number of elements copied which is min of sizes of both. 
        src and dst may be the same and range may overlap    
        a := []int{0,1,2,3,4,5,6,7,8,9}
        b := make([]int,len(a), 2*len(a))
        copy(b, a)


- type slice_split func([] int)([] int, [] int) :: type can be used to shorten function signature of a higher order function
- time package start := time.Now() time.Now().sub(start) or time.Since(start)

- ways to declare an array
    - [10]int{ 1, 2, 3 }
    - [...]int{5, 6, 7, 8, 22} :: compiler will count the number of elements in the ini.
    - [5]string{3: "Chris", 4: "Ron"}

- ways to declare a slice
    - var p []int :: gives you a slice without need to specify size
    - p := make([]int, len) // makes a slice with given len and cap = len

- slice2 := slice1[start:end] means copy over slice1 [start, end). Also, if start or end are omitted they imply 0, len()
    BEWARE : if a new slice is made out of an old array / another slice and then that old slice is changed, it reflect in new slice too.
    Also, if new slice was not including all elements of old slice towards it's end [don't care: len(old) - 2] then resizing new slice to inclrease it's
    length at end new = new[:len(new) + 1] will reveal the extra values in old slice that were earlier trimmed.

- passing slice / arr as reference aFunction(arr[:]). Also a function expecting a slice can be passed an array sliced by using [:]
- make slice and underlying array at once : slice1 := make([]type, len, cap) may skip len n cap. works the same as new([cap]type)[:len]
- make is applicable to slice , map , channel. new works with arrays and structs

- similar to for (x : coll) loop in java 
      for row := range 2dslice {
        for column := range 2dslice {

- Sorting
    - sort.Ints([]int)

- new() always returns a pointer to the type (primitive or user defined) that it is instantiating. though value can be accessed as x.key1 but to truly get the object use &x

- making changes to a string
    - to a position in string
        a := "str"
        a_chars := []rune(a)
        a_chars[0] = 'b'
        a = string(a_chars)
    - get a substring substr := str[start:end]

- fmt.Println("Array item", i, "is", a[i])
- math.MinInt32 similar to Integer.MIN_VALUE

- MAP:
    - freqmap := map[int]int
    - initializing with values :: map[string]int{"k1": 1, "k2": 2}
    - can store any kind of value e.g. functions too :: mf := map[int]func() int
    - access pattern : val1, isPresent = map1[key1] . This should be preferred as val1 = map[key1] will contain default value of the data type of value if the key is not present. _, ok pattern is go's way of expressing an error.
    - delete(map1, key1) deleting a key from map
    - keySet :: for key, value := range map1 { :: similar to how u access slices, instead of index, here we got keys

- structs
    - structs in go are declared as type Person struct {height, age int}
    - constructors are not possible in go, but factory function can be used ::
     func(height, age int) *Person {
         if height < 100 || age < 18 {
             return nil
         }

         return &Person{height, age}
    }
    By making person as not exported in this package but ths ff as exported, we can force other packages to use this ff.

    - read about anaonymous feilds in a struct in go
    - if two structs have similar structure one can be converter to another and comparision then checks all the values
    
    type a struct {
	    x, y, z int
    }
    type b struct {
        x, y, z int
    }
    func play() {
        ai := a{1,2,3}
        bi := b{1,2,3}
        fmt.Println(ai == a(bi)) // true
    }
    - structs inherit all the methods of any anonymous sub field. e.g.
    type Human struct {
        name string
        age int
    }
    type Student struct {
        gpa int
        Human
    }
    s1 := Student{gpa: 3.7, name: "Piyush", age: 23}. This pattern is similar to Interface

    - multiple inheritence smoked by golang
    type Camera struct { }

    func (c *Camera) TakeAPicture() string {  // method of Camera
        return "Click"
    }
    type Phone struct { }

    func (p *Phone ) Call() string {  // method of Phone
        return "Ring Ring"
    }

    // multiple inheritance
    type SmartPhone struct {    // can use methods of both structs
        Camera
        Phone
    }
- tags
    - type Person struct {
        Age int `json:"age"`
        Height int `json:"height" bson:"_height"`
    }

    these tags can be accessed by reflect package as follows :

	p := Person{Age: 23, Height: 175}

	if field, ok := reflect.TypeOf(p).FieldByName("Age"); ok {

		tag := field.Tag.Get("json")
		fmt.Println(tag)

	}

- methods
    - Any type can have methods, even a function type or alias types for int, bool, string, or array
    - receiver cannot be an interface type : since an interface is an abstract definition and a method is the implementation.
    - func (r *Rect) area int {
        return r.length * r.width
    }
    - method overloading:
        - like func, methods can't be overloaded (they are funcs)
        - however an overload a type based on another type (type apple fruit) is allowed. here we might want to replace formula to get volume from a generic one for all fruits
        - not allowed on an aliased type (type boundarylength = perimeter).
        - Methods are written independent of underlying type unlike in classes in oops where methods are written inside class.
        - If receiver type of methods are different (supposed to act on different type) then method names can be identical
        - func (c *Circle) calcArea float32 { return pi * math.Pow(c.radius, 2)} and func (s * Square) calcArea float32 {return s.side * s.side} are allowed as both have diff reeceiver type
        - methods cannot be defined on original data types of go. if needed, alias them first
    - getters and setters would be a good pattern so u don't have to export/expose data fields directly.
        - for setter use pattern p.setName(), for getter p.Name()
    - to format the printing of a struct String() can be overloaded
    func (tn *TwoInts) String() string {
            return "(" + strconv.Itoa(tn.a) + " / " + strconv.Itoa(tn.b) + ")"
        }

- GC
    package main
    import (
        "runtime"
        "time"
    )

    func main() {
        ms := runtime.MemStats{}
        runtime.ReadMemStats(&ms)
            
        println("Heap after GC. Used:", ms.HeapInuse, " Free:", ms.HeapIdle, " Meta:", ms.GCSys)

        time.Sleep(5 * time.Second)
    }

    - suppose some function needs to be run before garbage collector removes it:
        runtime.SetFinalizer(obj, func(obj *typeObj))

- Interface
    - checking if a value implements an interface
        type Stringer interface { String() string }
        if sv, ok := v.(Stringer); ok {
            fmt.Printf("v implements String(): %s\n", sv.String()); // note: sv, not v
        }