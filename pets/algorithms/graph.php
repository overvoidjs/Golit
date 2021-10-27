<?php
//Реализуем очередь FIFO
class Queue {
  // Массив очереди
  public $data = [];

  // Получаем длину очереди
  public function getLength():int {
    return count($this->data);
  }
  // Определяем, пуста ли очередь
  public function isEmpty():bool {
    return $this->getLength() === 0;
  }

  //Добавляем в конец очереди
  public function pushBack($element) {
    $this->data[] = $element;
  }

  //Получаем первый элемент очереди и удаляем его из очереди
  public function getFront():string {
    return (string)array_shift($this->data);
  }

  //Вернуть все
  public function getAll():array {
    return $this->data;
  }

}


function isManager(string $name):bool {
  //Тупо для теста, если имя заканчивается на м - то менеджер
  if( substr($name, -1) == "m" ){
    return true;
  } else {
    return false;
  }
}


function findManagerInGraph(array $graph):string {
  // инициализируем очередь
  $q = new Queue();

  //Добавляем в очередь первичный графф, связи 1 уровня
  foreach ($graph["you"] as $key => $value) {
    $q->pushBack($value);
  }

  //Работаем с очередью
  while ( !$q->isEmpty() ) {
    $name = $q->getFront();

    if( isManager($name) ){
      return $name;
    } else {
      if( count($graph[$name]) > 0 ){
        foreach ($graph[$name] as $key => $value) {
          $q->pushBack($value);
        }
      }
    }

  }

  return "Нету менеджеров";
}


$graph = [];
$graph["you"] = ["alice","bob","claire"];
$graph["bob"] = ["anuj","peggy"];
$graph["alice"] = ["peggy"];
$graph["claire"] = ["thom", "jonny"];
$graph["anuj"] = [];
$graph["peggy"] = [];
$graph["thom"] = [];
$graph["jonny"] = [];

var_dump(findManagerInGraph($graph));
