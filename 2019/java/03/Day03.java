import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;
import java.util.stream.Collectors;

public class Day03 {

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

        int size = 20000;
        int[][] grid = new int[size][size];

        String[] line1parts = lines.get(0).split(",");
        String[] line2parts = lines.get(1).split(",");
        int minDist = 99999;
        int minSteps = 99999;

        List<Instruction> instructions1 = Arrays.stream(line1parts)
            .map(part -> new Instruction(part)) // Map each part to an Instruction
            .collect(Collectors.toList()); // Collect the results into a List
        List<Instruction> instructions2 = Arrays.stream(line2parts)
            .map(part -> new Instruction(part)) // Map each part to an Instruction
            .collect(Collectors.toList()); // Collect the results into a List

        int[] pos = { 0, 0 };
        int posOffset = size / 2;
        int step = 0;

        for (int i = 0; i < instructions1.size(); i++) {
            Instruction instruction = instructions1.get(i);
            // execute each instruction
            for (int m = 0; m < instruction.mag; m++) {
                pos[0] += instruction.getDirVector()[0];
                pos[1] += instruction.getDirVector()[1];
                step++;
                grid[pos[0] + posOffset][pos[1] + posOffset] = step;
            }
        }

        pos = new int[] { 0, 0 };
        step = 0;

        for (int i = 0; i < instructions2.size(); i++) {
            Instruction instruction = instructions2.get(i);
            // execute each instruction
            for (int m = 0; m < instruction.mag; m++) {
                pos[0] += instruction.getDirVector()[0];
                pos[1] += instruction.getDirVector()[1];
                step++;
                // check for intersection
                if (grid[pos[0] + posOffset][pos[1] + posOffset] != 0) {
                    int dist = Math.abs(pos[0]) + Math.abs(pos[1]);
                    int totalSteps =
                        grid[pos[0] + posOffset][pos[1] + posOffset] + step;
                    if (dist < minDist) {
                        minDist = dist;
                    }
                    if (totalSteps < minSteps) {
                        minSteps = totalSteps;
                    }
                }
            }
        }

        // for (int r = 0; r < grid.length; r++) {
        //     for (int c = 0; c < grid[0].length; c++) {
        //         System.out.print(grid[r][c] ? "#" : "_");
        //     }
        //     System.out.println();
        // }
        System.out.println(minDist);
        System.out.println(minSteps);
    }
}
