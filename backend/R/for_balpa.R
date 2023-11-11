library(tidyverse)
library(rvest)
url <- 'https://www.basketball-reference.com/players/g/grantje01/gamelog-advanced/2024/'
print("ÖNÜ")

tryCatch(read_html(url))

print("test")

df <- df %>% select(Date, Tm, MP, TS., USG.)

df <- df %>% filter(MP != "Inactive")
df <- df %>% filter(MP != "Did Not Play")
df <- df %>% filter(MP != "MP")
df <- df %>% filter(MP != "Did Not Dress")

df$TS. <- as.numeric(df$TS.)
df$USG. <- as.numeric(df$USG.)
df$MP <- parse_number(df$MP)

df <- df %>%
  filter(MP >= 25)

df$TS. <- 100 * df$TS.

str(df)

# pl1 <- df %>% 
#   ggplot(aes(x = USG., y = TS.))  +
#   geom_point(aes(fill = MP), shape = 21, size = 4, alpha = .75) +  
#   geom_smooth(se = T, color = 'white', linetype = 'dashed') + 
#   theme(legend.position = 'none') + 
#   labs(x = "Usage Rate", 
#        y = "True Shooting %", 
#        title = "Jerami Grant's Skill Curve", 
#        subtitle = "Usage and scoring efficiency in games with at least 25 minutes played | 2023-24 Season") +
#   theme(plot.title.position = 'plot', 
#         plot.title = element_text(face = 'bold', size = 14, hjust=.5, vjust=-1), 
#         plot.subtitle = element_text(size = 10, hjust=.5, vjust=-1))

# ggsave("test.png", height=8, width=10, dpi="retina", plot=pl1)