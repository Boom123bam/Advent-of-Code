import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;
import java.util.stream.Collectors;

public class Day04 {

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

        int num1 = Integer.parseInt(lines.get(0).substring(0, 6));
        int num2 = Integer.parseInt(lines.get(0).substring(7));

        int count1 = 0;
        int count2 = 0;
        for (int i = num1; i <= num2; i++) {
            if (!pwHasAdjacentMatch(i) || !pwDigitsOnlyIncrease(i)) {
                continue;
            }
            count1++;
            if (pwHasExactly2AdjacentMatch(i)) {
                count2++;
            }
        }
        System.out.println(count1);
        System.out.println(count2);
    }

    public static boolean pwHasAdjacentMatch(int pw) {
        int last = pw % 10;
        while (pw > 0) {
            pw /= 10;
            if (pw % 10 == last) {
                return true;
            }
            last = pw % 10;
        }
        return false;
    }

    public static boolean pwDigitsOnlyIncrease(int pw) {
        int last = pw % 10;
        while (pw > 0) {
            pw /= 10;
            if (pw % 10 > last) {
                return false;
            }
            last = pw % 10;
        }
        return true;
    }

    public static boolean pwHasExactly2AdjacentMatch(int pw) {
        int last = pw % 10;
        int numSameDigit = 1;
        while (pw > 0) {
            pw /= 10;
            if (pw % 10 == last) {
                numSameDigit++;
            } else {
                if (numSameDigit == 2) {
                    return true;
                }
                numSameDigit = 1;
            }
            last = pw % 10;
        }
        if (numSameDigit == 2) {
            return true;
        }
        return false;
    }
}
