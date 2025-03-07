<?php
$start = microtime(true);

for ($i = 0; $i < 50000; $i++) {
    echo "Perulangan ke $i\n";
}

$duration = (microtime(true) - $start);
echo "Waktu eksekusi: " . $duration . " detik\n"; //12.225188016891 detik
