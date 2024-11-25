<?php

class HighSchoolSweetheart
{
    public function firstLetter(string $name): string
    {
        return $name[0];
    }

    public function initial(string $name): string
    {
       $ini = strtoupper($name[0]);
       return "$ini.";
    }

    public function initials(string $name): string
    {
        $arr = explode(" ", $name);
        $first = $this->initial($arr[0]);
        $sec = $this->initial($arr[1]);
        return "$first $sec";
    }

    public function pair(string $sweetheart_a, string $sweetheart_b): string
    {
        $name1 = $this->initials($sweetheart_a);
        $name2 = $this->initials($sweetheart_b);
        $ret = <<<END
     ******       ******
   **      **   **      **
 **         ** **         **
**            *            **
**                         **
**     $name1  +  $name2     **
 **                       **
   **                   **
     **               **
       **           **
         **       **
           **   **
             ***
              *
END;
        return $ret;
    }
}
