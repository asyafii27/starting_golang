let start = performance.now();

for (let i = 0; i < 50000; i++) {
    console.log("Perulangan ke", i);
}

let duration = (performance.now() - start) / 1000;
console.log("Waktu eksekusi:", duration.toFixed(6), "detik"); //  14.10102 detik
