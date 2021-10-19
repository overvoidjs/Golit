<?php

function getHalfOfSlice(array $arr): int {
  $x = count($arr) / 2;

  return (int)$x;
}

function qsort(array $arr): array {
  if(count($arr) <= 1){
    return $arr;
  } elseif (count($arr) == 2) {
    if($arr[0] > $arr[1]){
      $result_arr[] = $arr[1];
      $result_arr[] = $arr[0];

      return $result_arr;
    } else {
      return $arr;
    }
  } else {
    $half_index = getHalfOfSlice($arr);
    $pivot = $arr[$half_index];

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

    $result_arr = $less_than;
    $result_arr[] = $pivot;
    $result_arr = array_merge($result_arr,$more_than);

    return $result_arr;
  }
}

$arr = [64,13,22,16,4,5];

var_dump(qsort($arr));
