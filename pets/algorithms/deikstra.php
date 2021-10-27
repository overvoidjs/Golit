<?php

//Граф соседей
$graph["start"] = [];
$graph["start"]["a"] = 6;
$graph["start"]["b"] = 2;
$graph["a"] = [];
$graph["a"]["fin"] = 1;
$graph["b"] = [];
$graph["b"]["a"] = 3;
$graph["b"]["fin"] = 5;
$graph["fin"] = [];

//Таблица стоимостей
$costs = [];
$costs["a"] = 6;
$costs["b"] = 2;
$costs["fin"] = INF;

//Таблица родителей
$parents = [];
$parents["a"] = "start";
$parents["b"] = "start";
$parents["in"] = null;

//Таблица уже обработанный узлов графа
$processed = [];

//Функция поиска узла с наименьшей стоимостью среди необработанных
function findLowestCostNode($costs){

  $lowest_cost = INF;
  $lowest_cost_node = null;
  global $processed;

  foreach ($costs as $key => $cost) {
    if( $cost < $lowest_cost && !in_array($key, $processed) ){
      $lowest_cost = $cost;
      $lowest_cost_node = $key;
    }
  }
  return $lowest_cost_node;
}

$node = findLowestCostNode($costs);

while ( !is_null($node) ) {
  $cost = $costs[$node];
  $neighbors = $graph[$node];
  foreach (array_keys($neighbors) as $key => $n) {
    $new_cost = $cost + $neighbors[$n];
    if( $costs[$n] > $new_cost ){
      $costs[$n] = $new_cost;
      $parents[$n] = $node;
    }
  }

  $processed[] = $node;
  $node = findLowestCostNode($costs);
}

var_dump( $costs );
