<?php

function qsort(array $arr): array {
  if(count($arr) < 2){
    return $arr;
  } else {
    $half_index = count($arr) / 2;
    $pivot = $arr[(int)$half_index];

    $less_than = [];
    $more_than = [];

    foreach ($arr as $key => $value) {
      if($value < $pivot){
        $less_than[] = $value;
      } elseif ($value > $pivot) {
        $more_than[] = $value;
      }
    }

    $less_than = qsort($less_than);
    $more_than = qsort($more_than);

    return array_merge($less_than,array($pivot),$more_than);
  }
}

$arr = [64,13,22,16,4,5];

var_dump(qsort($arr));
