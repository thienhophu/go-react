# Build Docker at local

docker build -t go-react .

# Run Docker at local

docker run --rm -d --name go-react -p 3000:3000 go-react

# Stop Docker at Local

docker container stop go-react

# Deploy master branch to heroku

git push heroku master
