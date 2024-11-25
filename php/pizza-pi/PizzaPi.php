<?php

class PizzaPi
{
    public function calculateDoughRequirement($pizzas, $persons)
    {
        return $pizzas * (($persons * 20) + 200);
    }

    public function calculateSauceRequirement($pizzas, $can_volume)
    {
        return $pizzas * 125 / $can_volume;
    }

    public function calculateCheeseCubeCoverage($cheese_dimension, $thickness, $diameter)
    {
        return floor(($cheese_dimension**3) / ($thickness * M_PI * $diameter));
    }

    public function calculateLeftOverSlices($pizzas, $persons)
    {
        return ($pizzas * 8) % $persons; 
    }
}
