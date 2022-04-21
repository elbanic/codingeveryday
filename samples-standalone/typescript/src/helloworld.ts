
import {readFile} from 'fs'
import * as R from "ramda"

let message: string = "Hello World";

console.log(message);

let n: number = 1;
console.log(n);

n = 2;
console.log(n);

let a;
console.log(a);

a = 1;
console.log(a);

let b = 'b';
console.log(b);

let o: object = { name: 'Jack', age: 32 };
o = { first: 1, second: 2 };

console.log(o);

interface IPerson {
    name: string
    age: number
}

let good: IPerson = { name: 'Jack', age: 32 };
//let bed: IPerson = { name: 'Jack'};
console.log(good);

interface IPerson2 {
    name: string
    age: number
    etc?: boolean
}

let good2: IPerson2 = { name: 'Jack', age:32};
let good3: IPerson2 = { name: 'Jack', age:32, etc: true};

console.log(good2);
console.log(good3);

let tuple: any[] = [1, true, 'aa'];
console.log(tuple);


const array: number[] = R.range(1,10)
R.pipe(
    R.tap(n => console.log(n))
)(array)

R.tap(n => console.log(n))(array)
console.log(array)

const incNumber = R.pipe(
    R.tap(a => console.log('before inc:', a)),
    R.map(R.inc),
    R.tap(a => console.log('after inc:', a))
)

const newNumber = incNumber(array)