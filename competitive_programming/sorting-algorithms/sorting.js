function main () {
    let arr = [1,2,7,3,7,0,8,6,4];
    selectionSort(arr);
    console.log(arr.join())
};

function selectionSort(a = []) {

    for (let i = 0; i < a.length-1; i++) {

        let min = i;

        for (let j = i+1; j < a.length; j++) {
            if (a[min] > a[j])
                min = j;
        }

        console.log(`it: ${i} has min ${a[min]} at ${min}`, a.join());

        let temp = a[min];
        a[min] = a[i];
        a[i] = temp;
    }

}



// calling the main function below
main();
module.exports = main;