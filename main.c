#include <limits.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define uint unsigned int

#define PLAY_LOC(grid, play) grid[play.row-1][play.col-1]

typedef struct {
    uint row;
    uint col;
} Play;

void print_grid(char* grid[3][3]) {
    printf("- 1 2 3\n");
    for (int i = 0; i < 3; i++) {
        printf("%d ", i+1);
        for (int j = 0; j < 3; j++) {
            printf("%s ", grid[i][j]);
        }
        printf("\n");
    }
}


bool is_corner_piece(uint row, uint col) {
    return (
        (row == 1 && col == 1)
        || (row == 3 && col == 1)
        || (row == 1 && col == 3)
        || (row == 3 && col == 3)
    );
}


char* check_victory(char* grid[3][3]) {
    // lines
    for (size_t i = 0; i < 3; i++) {
        if(
            strcmp(grid[i][0], "□") == 0
            || strcmp(grid[i][1], "□") == 0
            || strcmp(grid[i][2], "□") == 0
        ){
            continue;
        }
        if (strcmp(grid[i][0], grid[i][1]) == 0 && strcmp(grid[i][1], grid[i][2]) == 0) {
            return grid[i][0];
        }
    }
    // cols
    for (size_t i = 0; i < 3; i++) {
        if(
            strcmp(grid[0][i], "□") == 0
            || strcmp(grid[1][i], "□") == 0
            || strcmp(grid[2][i], "□") == 0
        ){
            continue;
        }
        if (strcmp(grid[0][i], grid[1][i]) == 0 && strcmp(grid[1][i], grid[2][i]) == 0) {
            return grid[0][i];
        }
    }
    // diagonals
    if(
        strcmp(grid[0][0], "□") != 0
        && strcmp(grid[1][1], "□") != 0
        && strcmp(grid[2][2], "□") != 0
    ){
        if (strcmp(grid[0][0], grid[1][1]) == 0 && strcmp(grid[1][1], grid[2][2]) == 0) {
            return grid[0][0];
        }
    }
        if(
        strcmp(grid[2][0], "□") != 0
        && strcmp(grid[1][1], "□") != 0
        && strcmp(grid[0][2], "□") != 0
    ){
        if (strcmp(grid[2][0], grid[1][1]) == 0 && strcmp(grid[1][1], grid[0][2]) == 0) {
            return grid[2][0];
        }
    }
    return NULL;
}

bool is_blocking_player_from_win(Play play, char* enemy_symbol, char* grid[3][3]) {
    PLAY_LOC(grid, play) = enemy_symbol;
    char *win_symbol = check_victory(grid);
    PLAY_LOC(grid, play) = "□";
    if (win_symbol == NULL) return false;
    return true;
}

bool is_in_path_to_win(Play play, char* player_symbol, char* grid[3][3]) {
    return false;
}

void generate_all_plays(char* play_symbol, char* grid[3][3], Play* plays) {
    size_t play_index = 0;
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 3; j++) {
            if (strcmp(grid[i][j], "□") == 0) {
                Play play = {.row = i+1, .col = j+1};
                plays[play_index] = play;
                play_index++;
            }
        }
    }
}

bool is_gonna_win(Play play, char* play_symbol, char* grid[3][3]){
    PLAY_LOC(grid, play) = play_symbol;
    char *win_symbol = check_victory(grid);
    PLAY_LOC(grid, play) = "□";
    if (win_symbol == NULL) return false;
    return true;
}

int calculate_play_cost(Play play, char* play_symbol, char* enemy_symbol, char* grid[3][3]) {
    // Cost considerations:
    // corner piece: - cost
    // blocking a player from win: - cost
    // play in path of win: - cost

    int cost = 1000;

    if (is_gonna_win(play, play_symbol, grid)) {
        return 0;
    }

    if (is_corner_piece(play.row, play.col)) {
        cost -= 100;
    }

    if (is_blocking_player_from_win(play, enemy_symbol, grid)) {
        cost -= 500;
    }

    if (is_in_path_to_win(play, play_symbol, grid)) {
        cost -= 250;
    }

    return cost;
}

bool play_once(char* play_symbol, Play play, char* grid[3][3]) {
    if (play.row > 3 || play.col > 3) {
        fprintf(stderr, "[\x1b[31mERROR\033[0m] row and col must be between 1 and 3\n");
        return false;
    }
    if (strcmp(PLAY_LOC(grid, play), "□") != 0) {
        fprintf(stderr, "[\x1b[31mERROR\033[0m] Your play is in an invalid square\n");
        return false;
    } else {
        PLAY_LOC(grid, play) = play_symbol;
    }
    return true;
}

void ai_play(char* play_symbol, char* enemy_symbol, char* grid[3][3]) {
    // consider all possible plays and minimize play cost

    Play plays[9] = {};
    int play_cost[9] = {};

    generate_all_plays(play_symbol, grid, plays);

    for (size_t i = 0; i < 9; i++) {
        Play play = plays[i];
        if (play.row == 0 && play.col == 0) {
            play_cost[i] = INT_MAX;
            continue;
        }
        play_cost[i] = calculate_play_cost(play, play_symbol, enemy_symbol, grid);
    }

    int min_cost = INT_MAX;
    Play best_play;
    for (size_t i = 0; i < 9; i++) {
        if (play_cost[i] < min_cost) {
            min_cost = play_cost[i];
            best_play = plays[i];
        }
    }

    play_once(play_symbol, best_play, grid);

}


bool player_turn(char* play_symbol, char* grid[3][3]) {
    printf("Your turn (row,col)\n");
    print_grid(grid);
    Play play;
    scanf("%i,%i", &play.row, &play.col);
    return play_once(play_symbol, play, grid);
}

void ai_turn(char* play_symbol,char* enemy_symbol, char* grid[3][3]) {
    printf("AI's turn\n");;
    ai_play(play_symbol, enemy_symbol, grid);
}

int main() {
    char* grid[3][3] = {
        {"□", "□", "□"},
        {"□", "□", "□"},
        {"□", "□", "□"}
    };

    char* choice_symbol;
    uint choice;
    char* choices[] = {"X", "O"};

    while(true) {
        printf("Choose your character:\n 1 -> X (game starter)\n 2 -> O\n");
        int read = scanf("%i", &choice);
        if (read != 1) {
            int c;
            while ((c = getchar()) != '\n' && c != EOF);
        }
        if (choice != 1 && choice != 2) {
            printf("[\x1b[31mERROR\033[0m] Must either provide 1 or 2!\n");
            continue;
        }
        break;
    }

    char* player_symbol = choices[choice-1];
    char* ai_symbol = choices[!(choice-1)];

    printf("You choose: %s\n", player_symbol);

    uint row, col;
    size_t total_plays = 0;
    if (choice == 1) {
        while (!player_turn(player_symbol, grid)){}
    } else {
        ai_turn(ai_symbol, player_symbol, grid);
    }
    total_plays++;

    int turn = (choice == 2);
    do {
        if (turn) {
            while (!player_turn(player_symbol, grid)){}
            turn = 0;
            total_plays++;
            continue;
        }
        ai_turn(ai_symbol, player_symbol, grid);
        turn = 1;
        total_plays++;
    } while(check_victory(grid) == NULL && total_plays < 9);

    print_grid(grid);

    if (total_plays == 9) {
        printf("It's a tie!\n");
        return 0;
    }

    printf("Good job, %s won!\n", check_victory(grid));
}
