<?php
/*
Изображение графа в этом примере

          a
     (6)/  \(1)
start     |    fin
     (2)\  /(5)
         b

Взят из книги Грокаем Алгоритмы. 7 Глава
*/

//Граф соседей
$graph["start"] = [];
$graph["start"]["a"] = 5;
$graph["start"]["b"] = 2;

$graph["a"] = [];
$graph["a"]["c"] = 4;
$graph["a"]["d"] = 2;
// $graph["a"]["fin"] = 1;
$graph["b"] = [];
$graph["b"]["a"] = 8;
$graph["b"]["d"] = 7;

$graph["c"] = [];
$graph["c"]["d"] = 6;
$graph["c"]["fin"] = 3;

$graph["d"] = [];
$graph["d"]["fin"] = 1;

$graph["fin"] = [];

//Таблица стоимостей
$costs = [];
$costs["a"] = 5;
$costs["b"] = 2;
$costs["c"] = INF;
$costs["d"] = INF;
$costs["fin"] = INF;

//Таблица родителей
$parents = [];
$parents["a"] = "start";
$parents["b"] = "start";
$parents["c"] = null;
$parents["d"] = null;

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

$node = findLowestCostNode($costs); //Находим узел с наименьшей стоимостью

while ( !is_null($node) ) { //Пока не обработаны все узлы
  $cost = $costs[$node]; //Текущая цена узла
  $neighbors = $graph[$node]; //Соседи этого узла

  foreach (array_keys($neighbors) as $key => $n) { //Пройдемся по всем соседям
    $new_cost = $cost + $neighbors[$n];
    if( $costs[$n] > $new_cost ){ //Если можно добраться быстрее
      $costs[$n] = $new_cost; //- обновим стоимость
      $parents[$n] = $node; // Назначим нового родитея
    }
  }

  $processed[] = $node;
  $node = findLowestCostNode($costs);
}


//Функция поиска оптимально маршрута по графу
function findOptimalWay($parents){
  $parents = array_reverse($parents, true);
  $first_key = array_keys($parents)[0];
  $way[] = $first_key;
  $way[] = $parents[$first_key];


  foreach ($parents as $key => $value) {
    if( isset($parents[$value]) && !in_array($parents[$value], $way) ){
      $way[] = $parents[$value];
    }
  }

  return array_reverse($way);
}

$x = findOptimalWay($parents);
var_dump( $x );
var_dump( $costs["fin"] );
