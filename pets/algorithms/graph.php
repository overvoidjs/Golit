<?php
//Реализуем очередь FIFO
class Queue {
  // Массив очереди
  public $data = [];

  // Получаем длину очереди
  public function getLength() {
    return count($this->data);
  }
  // Определяем, пуста ли очередь
  public function isEmpty() {
    return $this->getLength() === 0;
  }

  //Добавляем в конец очереди
  public function pushBack($element) {
    $this->data[] = $element;
  }

  //Получаем первый элемент очереди и удаляем его из очереди
  public function getFront() {
    return array_shift($this->data);
  }

  public function getAll(){
    return $this->data;
  }

}


$q = new Queue();

$q->pushBack('a');
$q->pushBack('b');
$q->pushBack('c');

var_dump($q->getFront());
var_dump($q->getAll());
