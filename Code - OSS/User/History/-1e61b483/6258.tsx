interface ProductComponentsProps{
    type: string,
    id: string | undefined
}

const desc = ("Camisa Tech Feminina - Leve, respirável e de secagem rápida, com modelagem ajustada que valoriza o corpo e garante conforto em qualquer ocasião." +
"Características: Tecido tecnológico de alta performance; Secagem rápida e respirabilidade; Leve, elástico e confortável;" +
"Modelagem feminina levemente ajustada;Ideal para treinos, viagens e uso diário")


export default function ProductComponents(data : ProductComponentsProps) {

    var title: string = "";
    var first_part: string = "";
    var second_title: string = "";
    var second_title: string = "";


    if (desc.indexOf('-') !== -1)
        title = desc.slice(0, desc.indexOf('-'));

    if (desc.indexOf('.') !== -1){
        first_part = desc.slice( ((title != "") ? desc.indexOf('-') : 0) , desc.indexOf('.'))
        first_part = desc.slice( ((title != '') ? desc.indexOf('-') : 0) , desc.indexOf('.'))
        second_title = desc.slice( ((second_title != "") ? desc.indexOf(":") : desc.indexOf('.') + 1) , desc.length)
    }

    if (data.type == 'DSC') {
        return (
            <div className="w-75 text-sm font-noto">
                <p>
                    <span className="font-bold">
                        {title}
                    </span>
                    {first_part}
                </p>

                <p className="mt-2">
                    <span className="font-bold">
                        {second_title}
                    </span>
                    {second_title}
                </p>
            </div>
        )
    }

    return (
        <div>
        </div>
    )
}