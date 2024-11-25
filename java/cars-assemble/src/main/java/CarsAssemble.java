public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        double ret = speed * 221;
        if (speed >= 1 && speed <= 4 ){
            return ret;
        } else if(speed >= 5 && speed <= 8 ){
            return ret * 0.9;
        }else if(speed == 9){
            return ret * 0.8 ;
        }else if(speed == 10 ){
            return ret * 0.77;
        }
        return 0;
    }

    public int workingItemsPerMinute(int speed) {
        return (int)(productionRatePerHour(speed) / 60);
    }
}
