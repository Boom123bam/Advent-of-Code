import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

public class Day01 {

    public static void main(String[] args) {
        List<String> lines = new ArrayList<String>();
        String filePath = "input.txt";
        try (BufferedReader br = new BufferedReader(new FileReader(filePath))) {
            String line;
            while ((line = br.readLine()) != null) {
                lines.add(line);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }

        int sum1 = 0;
        int sum2 = 0;
        for (int i = 0; i < lines.size(); i++) {
            try {
                int num = Integer.parseInt(lines.get(i));
                sum1 += num / 3 - 2;
                sum2 += getFuel2(num) - num;
            } catch (NumberFormatException e) {
                System.out.println("oops");
            }
        }
        System.out.println(sum1);
        System.out.println(sum2);
    }

    public static int getFuel2(int n) {
        int fuel = (n / 3) - 2;
        return fuel > 0 ? n + getFuel2(fuel) : n;
    }
}
