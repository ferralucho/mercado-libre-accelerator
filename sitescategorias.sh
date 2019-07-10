urlmercadolibre='https://api.mercadolibre.com'
urlsites=${urlmercadolibre}'/sites/'
urlcategorias='/categories'

function meli_sites_categories () {

local sites=$(curl -X GET ${urlsites}${urlcategorias})
local sitios=$(curl -s "https://api.mercadolibre.com/sites#json" | sort)
local nombresitios=$(curl -s "https://api.mercadolibre.com/sites#json" | jq -r '.[].name | tostring' | sort )

nombresitio=$1
categorianombre=$2

if [ "$#" -eq 0 ]; then
    echo 'Elija un país: '
    options=($nombresitios)
    select opt in "${options[@]}"
    do
        case $opt in $opt)       

    paisseleccionado=$(echo "$opt")
    sitio=$(echo ${sitios} | jq --arg paisseleccionado $paisseleccionado -c 'map(select(.name == $paisseleccionado))')                            
    idsitio=$(echo "$sitio" | jq '.[] | .id' | sed -e 's/^"//' -e 's/"$//')   
    categorias=$(curl -s -X GET ${urlsites}${idsitio}${urlcategorias})    
                                        
                echo -e ' \t '${paisseleccionado}
                echo -e ' \t ' '------------'
                echo -e ' \t ' "$categorias" | jq -r '.[].name'           
                ;;       
        esac
    done
else
    sitio=$(echo ${sitios} | jq --arg nombresitio $nombresitio -c 'map(select(.name == $nombresitio))') 
    idsitio=$(echo "$sitio" | jq '.[] | .id' | sed -e 's/^"//' -e 's/"$//')    
    categorias=$(curl -s -X GET ${urlsites}${idsitio}${urlcategorias})  
     if [ "$categorianombre" ]
                then
                    echo -e ' \t ' ${nombresitio}
                    echo '------------'
                    categoriaencontrada=$(echo "${categorias}" | jq --arg categorianombre $categorianombre -c 'map(select(.name | contains($categorianombre)))' | jq '.[] | .name' )
                    if [ "$categoriaencontrada" ]; then
                        echo 'Se encontró ' 
                        echo "$categoriaencontrada"                                          
                    else
                        echo 'No existe tal categoría para el país seleccionado.'
                    fi
                else
                    echo ${nombresitio}
                    echo '------------'
                    echo "${categorias}" | jq -r '.[].name' 
     fi
fi
}
meli_sites_categories $1 $2
