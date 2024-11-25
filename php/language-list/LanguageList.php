<?php

function language_list(...$items)
{
    return $items;
}

function add_to_language_list($items, $new){
    array_push($items, $new);
    return $items;
}


function prune_language_list($items){
    array_shift($items);
    return $items;
}

function current_language($items){
    return $items[0];
}

function language_list_length($items){
    return count($items);
}