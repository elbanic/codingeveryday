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